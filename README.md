# cofiswarm-infer-vllm

Cofiswarm component: `infer-vllm`.

- Layout: [REPO-STANDARD-LAYOUT](https://github.com/keepdevops/cofiswarm-docs/blob/main/REPO-STANDARD-LAYOUT.md)
- Migration: [MIGRATION-SPRINTS](https://github.com/keepdevops/cofiswarm-docs/blob/main/MIGRATION-SPRINTS.md)

## FHS paths

| Path | Purpose |
|------|---------|
| `/etc/cofiswarm/infer-vllm/` | config |
| `/var/lib/cofiswarm/infer-vllm/` | state |
| `/var/log/cofiswarm/infer-vllm/` | logs |

## Test

```bash
./test/scripts/assert-layout.sh infer-vllm
```
