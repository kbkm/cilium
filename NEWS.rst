****
NEWS
****

HEAD
====

Fixes
-----

- k8s NetworkPolicy fixes
  - Correctly handle k8s NetworkPolicy matchLabels
  - Allow all sources if []NetworkPolicyPeer is empty or missing

Features
--------

- Improved logging readability (`GH #499 <https://github.com/cilium/cilium/pull/499>`_)
- Reduced size of cilium binary from 27M to 17M
- Decreased endpoint operations time by introducing parallelization in regeneration
- Replaced all endpoint synchronous CLI operations with asynchronous CLI operations
- Allow rule now supports matching multiple labels
- Kubernetes NetworkPolicy improvements
  - Support L4 filtering with v1beta1.NetworkPolicyPort

0.8.0
=====

- First initial release
