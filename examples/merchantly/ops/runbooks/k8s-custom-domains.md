---
title: K8s Custom Domains — SRE Runbook
doc_kind: ops
doc_function: runbook
purpose: Операционные процедуры для управления кастомными доменами в Kubernetes. Читать при инцидентах с доменами.
derived_from:
  - ../../../../memory-bank/dna/governance.md
  - ../production.md
status: active
audience: humans_and_agents
---

# K8s Custom Domains — SRE Runbook

## Quick Links
- **Dashboard:** [Grafana K8s Custom Domains](https://grafana.example.com/d/k8s-custom-domains)
- **Alert Manager:** [K8s Custom Domains Alerts](https://alertmanager.example.com/)
- **Prometheus:** [K8s Custom Domains Metrics](https://prometheus.example.com/graph?g0.expr=k8s_)

---

## 1. Too Many Pending Custom Domains

**Alert:** `TooManyPendingCustomDomainsWarning` (>10 for 15m) | `PendingDomainsExceedThreshold` (>20 for 5m)

### Symptoms
- Grafana shows spike in "Pending Custom Domains" gauge
- Users report domains not becoming active
- `certificate_ready_time_seconds` high

### Root Causes
1. **cert-manager not issuing certificates** — Check cert-manager pod logs
2. **ACME challenge failing** — DNS/HTTP validation issue
3. **CheckAllPendingCertificatesJob stuck** — Job queue backlog
4. **K8s API overload** — Too many concurrent API calls

### Diagnosis

```bash
# Check pending vendors in DB
psql $DATABASE_URL << 'EOF'
SELECT id, domain, ingress_status, ingress_error, ingress_updated_at
FROM vendors
WHERE ingress_status = 'pending'
AND ingress_updated_at < NOW() - INTERVAL '10 minutes'
ORDER BY ingress_updated_at ASC;
EOF

# Check K8s certificates status
kubectl -n merchantly get cert -l app.kubernetes.io/managed-by=merchantly

# Check cert-manager logs
kubectl -n cert-manager logs -l app=cert-manager --tail=100

# Check Ingress HTTP-01 challenge
kubectl -n merchantly get ingress -l app.kubernetes.io/managed-by=merchantly
kubectl -n merchantly describe ingress custom-domain-example-com
```

### Resolution

#### Option A: Restart cert-manager (if stuck)
```bash
kubectl -n cert-manager rollout restart deployment/cert-manager
kubectl -n cert-manager rollout restart deployment/cert-manager-webhook
```

#### Option B: Force certificate reconciliation
```bash
# Delete stuck certificate to trigger re-issuance
kubectl -n merchantly delete cert custom-domain-example-com-tls --cascade=orphan

# Or trigger manual reconciliation
kubectl -n merchantly exec deployment/app -- bundle exec rails runner \
  'K8sReconciliationJob.perform_now'
```

#### Option C: Check CheckAllPendingCertificatesJob
```bash
# Check if job is running
psql $DATABASE_URL << 'EOF'
SELECT * FROM solid_queue_jobs
WHERE class_name = 'CheckAllPendingCertificatesJob'
ORDER BY created_at DESC LIMIT 5;
EOF

# If stuck, restart SolidQueue workers
kubectl -n merchantly rollout restart deployment/app
```

### Prevention
- Monitor `certificate_ready_time_seconds` p95 threshold
- Keep cert-manager updated
- Monitor ACME rate limits (Let's Encrypt)
- Ensure DNS records are properly pointed to cluster

---

## 2. Ingress Creation Failure Rate High

**Alert:** `IngressCreationFailureRate` (>5% for 5m)

### Symptoms
- Many vendors stuck in "failed" status
- `ingress_error` = `INGRESS_CREATE_FAILED` or `K8S_UNAVAILABLE`
- K8s API call duration spiking

### Root Causes
1. **K8s API unavailable/slow** — Cluster issues
2. **RBAC permissions missing** — ServiceAccount not configured
3. **Label validation failure** — Label conflicts
4. **Duplicate Ingress** — Domain_conflict error

### Diagnosis

```bash
# Check K8s API health
kubectl cluster-info

# Check Merchantly app RBAC
kubectl -n merchantly auth can-i create ingresses --as=system:serviceaccount:merchantly:merchantly-domain-manager

# Check app logs for K8s errors
kubectl -n merchantly logs -l app=app --tail=200 | grep -i "k8s\|ingress\|error"

# Check if any Ingress with conflict
kubectl -n merchantly get ingress -o yaml | grep -i "domain\|vendor"
```

### Resolution

#### Option A: Verify RBAC
```bash
# Check if role binding exists
kubectl -n merchantly get rolebinding | grep ingress

# Reapply RBAC manifests
kubectl apply -f k8s/manifests/k8s-custom-domains-rbac.yaml
```

#### Option B: Fix K8s API issues
```bash
# Check K8s API server status
kubectl get cs

# If APIServer not healthy, contact K8s cluster admin
# Escalate to infrastructure team
```

#### Option C: Clean up failed vendors
```bash
# Identify failed vendors
psql $DATABASE_URL << 'EOF'
SELECT id, domain, ingress_error, ingress_updated_at
FROM vendors
WHERE ingress_status = 'failed'
AND ingress_error = 'INGRESS_CREATE_FAILED';
EOF

# For each vendor, either:
# 1. Retry manually
psql $DATABASE_URL << 'EOF'
UPDATE vendors SET ingress_status = NULL, ingress_error = NULL
WHERE id = 123;
EOF

# Then enqueue new job
bin/rails runner 'CustomDomainIngressJob.perform_later(123, :install, "shop.example.com", Time.current)'

# 2. Or check domain conflict
kubectl -n merchantly get ingress custom-domain-shop-example-com -o yaml
```

---

## 3. Orphaned Ingress Resources

**Alert:** `OrphanedIngressResourcesPresent` (>0 for 5m)

### Symptoms
- `k8s_orphaned_ingress_total` gauge shows > 0
- Ingress exists in K8s but no matching vendor in DB
- Likely from failed archive operations

### Root Causes
1. **Archive job never executed** — Job queue backlog
2. **Archive job failed silently** — Exception in async job
3. **Vendor deleted before Ingress cleanup** — Race condition
4. **Database restoration without K8s cleanup** — Restore operation

### Diagnosis

```bash
# List all custom domain Ingress
kubectl -n merchantly get ingress -l app.kubernetes.io/managed-by=merchantly

# Check each Ingress labels
kubectl -n merchantly get ingress custom-domain-shop-example-com -o jsonpath='{.metadata.labels}'

# Check if vendor exists in DB
psql $DATABASE_URL << 'EOF'
SELECT id, domain, archived_at
FROM vendors
WHERE id = (SELECT label from ingress where name = 'custom-domain-shop-example-com')
EOF

# If vendor not found or archived, Ingress is orphaned
```

### Resolution

#### Option A: Delete orphaned Ingress
```bash
# Single Ingress
kubectl -n merchantly delete ingress custom-domain-shop-example-com

# All orphaned (CAUTION: verify first)
kubectl -n merchantly delete ingress -l app.kubernetes.io/managed-by=merchantly \
  --field-selector metadata.name!=custom-domain-* # adjust as needed
```

#### Option B: Restore vendor if accidentally deleted
```bash
# If vendor should exist but was deleted, restore from backup
# Contact DBA to restore from automated backups
```

#### Option C: Trigger reconciliation to auto-cleanup
```bash
bin/rails runner 'K8sReconciliationJob.perform_now'
```

### Prevention
- Ensure `K8sReconciliationJob` runs hourly (check SolidQueue)
- Test archive flow in staging before deploying
- Monitor `k8s_reconciliation_operations_total` counter

---

## 4. Certificate Timeout

**Alert:** `CertificateReadyTimeIncreasing` (p95 > 20min for 10m)

### Symptoms
- Vendors stuck in "pending" with `ingress_status = 'pending'` > 20 minutes
- `certificate_ready_time_seconds` histogram shows high latency
- `ingress_error = 'CERT_TIMEOUT'`

### Root Causes
1. **ACME challenge not resolving** — DNS not pointing to cluster
2. **cert-manager webhook issues** — K8s webhook not responding
3. **Let's Encrypt rate limit** — Too many requests
4. **Network/firewall blocking HTTP-01 challenge**

### Diagnosis

```bash
# Check certificate status in K8s
kubectl -n merchantly get cert custom-domain-shop-example-com -o yaml

# Check cert-manager logs for ACME errors
kubectl -n cert-manager logs -l app=cert-manager | grep -i "acme\|challenge"

# Manually test HTTP-01 challenge
curl -v http://shop.example.com/.well-known/acme-challenge/test-token

# Check DNS resolution
nslookup shop.example.com
dig shop.example.com

# Check Let's Encrypt rate limits
curl https://letsencrypt.org/stats/
```

### Resolution

#### Option A: Fix DNS/network
```bash
# Ensure domain points to K8s cluster ingress IP
nslookup shop.example.com
# Should resolve to K8s ingress IP (e.g., 1.2.3.4)

# Verify with dig
dig shop.example.com +short

# If wrong, update DNS records and wait for propagation (5-10 min)
```

#### Option B: Restart cert-manager
```bash
kubectl -n cert-manager rollout restart deployment/cert-manager
kubectl -n cert-manager rollout restart deployment/cert-manager-webhook
```

#### Option C: Delete and recreate certificate
```bash
kubectl -n merchantly delete cert custom-domain-shop-example-com --cascade=orphan
sleep 10

# Trigger new certificate creation
bin/rails runner 'vendor = Vendor.find(123); \
  CustomDomainIngressJob.perform_later(vendor.id, :install, vendor.domain, Time.current)'
```

#### Option D: Check Let's Encrypt rate limits
```bash
# If rate-limited, wait before retrying
# Rate limits reset daily at midnight UTC
# See: https://letsencrypt.org/docs/rate-limits/
```

### Prevention
- Monitor `certificate_ready_time_seconds` percentiles
- Pre-verify DNS records before users set custom domain
- Document Let's Encrypt rate limits in user docs
- Consider using wildcard certificates for high-volume vendors

---

## 5. Failed Custom Domains (ingress_error)

**Alert:** `FailedCustomDomainsPresent` (>0 for 5m)

### Symptoms
- Vendors with `ingress_status = 'failed'`
- `ingress_error` contains error code (e.g., `DOMAIN_CONFLICT`, `K8S_UNAVAILABLE`)

### Common Error Codes

| Error Code | Meaning | Fix |
|------------|---------|-----|
| `DOMAIN_CONFLICT` | Domain used by another vendor | Contact support, reassign domain |
| `CERT_TIMEOUT` | Certificate not ready in 20m | See "Certificate Timeout" section |
| `K8S_UNAVAILABLE` | K8s API unreachable | Check K8s cluster health |
| `INGRESS_CREATE_FAILED` | Ingress creation failed | Check RBAC, K8s capacity |
| `CERT_FAILED` | cert-manager reported failure | Check cert-manager logs |

### Resolution

#### For DOMAIN_CONFLICT
```bash
# Find which vendor owns the domain
psql $DATABASE_URL << 'EOF'
SELECT id, name, domain FROM vendors WHERE domain = 'shop.example.com';
EOF

# Option 1: Unassign domain from one vendor
UPDATE vendors SET domain = NULL WHERE id = 456;

# Option 2: Update Ingress label to correct owner
kubectl -n merchantly patch ingress custom-domain-shop-example-com \
  --type json -p '[{"op":"replace","path":"/metadata/labels/merchantly.io~1vendor-id","value":"123"}]'

# Retry job
bin/rails runner 'CustomDomainIngressJob.perform_later(123, :install, "shop.example.com", Time.current)'
```

#### For K8S_UNAVAILABLE
```bash
# Check K8s cluster
kubectl cluster-info

# If cluster down, escalate to infrastructure team
# Otherwise, retry job after 5 minutes
bin/rails runner 'CustomDomainIngressJob.perform_later(123, :install, "shop.example.com", Time.current)'
```

#### For INGRESS_CREATE_FAILED
```bash
# Check app logs
kubectl -n merchantly logs -l app=app --tail=100 | grep -i ingress

# Verify RBAC
kubectl -n merchantly auth can-i create ingresses --as=system:serviceaccount:merchantly:merchantly-domain-manager

# Retry
bin/rails runner 'CustomDomainIngressJob.perform_later(123, :install, "shop.example.com", Time.current)'
```

---

## 6. Reconciliation Job Not Running

**Alert:** `ReconciliationJobNotRunning` (no activity for 2 hours)

### Symptoms
- No `k8s_reconciliation_operations_total` counter increments
- Orphaned Ingress accumulating
- SolidQueue not picking up recurring jobs

### Root Causes
1. **SolidQueue workers not running** — Deployment issue
2. **Recurring job not registered** — Database corruption
3. **SolidQueue database unavailable** — DB connectivity issue

### Diagnosis

```bash
# Check SolidQueue pod status
kubectl -n merchantly get pods -l app=app | grep solid-queue
kubectl -n merchantly logs -l app=app,component=queue --tail=50

# Check if recurring job registered
psql $DATABASE_URL << 'EOF'
SELECT * FROM solid_queue_scheduled_executions
WHERE class_name = 'K8sReconciliationJob';
EOF

# Check SolidQueue database connectivity
psql $DATABASE_URL -c "SELECT COUNT(*) FROM solid_queue_jobs;"
```

### Resolution

#### Option A: Restart SolidQueue workers
```bash
kubectl -n merchantly rollout restart deployment/app
# or
systemctl restart solid_queue.service
```

#### Option B: Re-register recurring job
```bash
bin/rails solid_queue:install
```

#### Option C: Manually trigger reconciliation
```bash
bin/rails runner 'K8sReconciliationJob.perform_now'
```

### Prevention
- Monitor `solid_queue_jobs` table for stuck jobs
- Ensure SolidQueue workers have separate processes (prod setup)
- Monitor `k8s_reconciliation_operations_total` metric in Prometheus

---

## 7. Monitoring Checklist

### Daily (check dashboard)
- [ ] Pending vendors gauge < 5
- [ ] Failed vendors gauge = 0
- [ ] Orphaned Ingress = 0
- [ ] Ingress creation success rate > 95%

### Weekly
- [ ] Certificate ready time p95 < 5 minutes
- [ ] No K8s API errors in last 7 days
- [ ] Reconciliation job ran hourly (at least 168 times)

### Monthly
- [ ] Review alert thresholds against actual distribution
- [ ] Audit ingress labels and ownership
- [ ] Check cert-manager quota usage
- [ ] Review Let's Encrypt rate limit usage

---

## 8. Emergency Procedures

### Disable Feature (rollback)
```bash
# In application.yml
feature_k8s_custom_domains: false

# Redeploy
make deploy

# This stops new Ingress creation, leaves existing ones in place
```

### Full Cleanup (nuclear option)
```bash
# Delete all custom domain Ingress
kubectl -n merchantly delete ingress -l app.kubernetes.io/managed-by=merchantly

# Reset vendor statuses
psql $DATABASE_URL << 'EOF'
UPDATE vendors
SET ingress_status = NULL, ingress_error = NULL, ingress_updated_at = NULL
WHERE k8s_custom_domains = true;
EOF

# Rebuild from scratch
bin/rails runner 'K8sReconciliationJob.perform_now'
```

### Report Issue
```bash
# Gather debug info
mkdir -p /tmp/k8s-debug
kubectl -n merchantly describe all > /tmp/k8s-debug/describe.txt
kubectl -n cert-manager logs --all-containers --tail=100 > /tmp/k8s-debug/cert-manager.log
psql $DATABASE_URL -c "SELECT * FROM vendors WHERE k8s_custom_domains = true LIMIT 10;" > /tmp/k8s-debug/vendors.txt

# Send to Bugsnag/Slack
# Include: /tmp/k8s-debug/* + alert name + time
```

---

## Useful Commands

```bash
# Get all custom domain resources
kubectl -n merchantly get all -l app.kubernetes.io/managed-by=merchantly

# Watch Ingress
kubectl -n merchantly get ingress -w

# Watch certificates
kubectl -n merchantly get cert -w

# Real-time logs
kubectl -n merchantly logs -f deployment/app | grep -i "k8s\|ingress\|cert"

# Check metrics
curl http://prometheus.example.com/api/v1/query?query=k8s_pending_vendors_gauge
```
