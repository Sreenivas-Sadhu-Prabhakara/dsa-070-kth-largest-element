# Kth Largest Element (Heap)

> Striver SDE Sheet 070 · Day 12 Heaps — a min-heap holding only the k largest; whatever sits at its root is the k-th largest. Go, stdlib only, tested.

## Problem

Find the k-th largest element of an array (duplicates count separately). This slot on the sheet wants the heap discipline; the quickselect version appears again at #150 with average-O(n) — same problem, different tool, deliberately.

## Approach

Keep a **min**-heap of capacity k — min, not max, because the heap's job is guarding the *smallest of the current top k*, which is exactly the element a new candidate must beat. Values ≤ the root are ignored; a larger value replaces the root (`heap.Fix` instead of pop-then-push does it in one sift). After the scan the root is the k-th largest. O(n log k) beats full sorting whenever k ≪ n, and the stream never needs to be stored — which is why this exact shape returns as "k-th largest in a stream" (#148).

**Complexity:** O(n log k) time, O(k) space.

## Run

    go test ./...   # verify
    go run .        # demo

---

One of 191 repos solving Striver's SDE Sheet in Go — stdlib only, one problem per repo, every solution tested. Browse the set: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
