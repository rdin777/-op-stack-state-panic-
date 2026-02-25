# -op-stack-state-panic-
Proof of Concept for a Nil Pointer Dereference vulnerability in Optimism-Geth (op-geth) state management, leading to node panic (DoS).
Markdown
# Vulnerability Research: Node Panic in Optimism-Geth (op-geth)

## Overview
This repository contains a Proof of Concept (PoC) for a **Nil Pointer Dereference** vulnerability discovered in the state management logic of `op-geth` (the engine powering Base and Optimism L2s). The vulnerability leads to an immediate node crash (DoS) during specific state revert operations.

## Technical Analysis
The panic occurs when the state journal attempts to revert a `nonceChange` for an account that has been destroyed and removed from the active state database within the same transaction/block.

- **Type:** Denial of Service (DoS)
- **Component:** `core/state`
- **Root Cause:** `nil pointer dereference` in `stateObject.setNonce`
- **Status:** Reported via HackerOne (Informative)

### Stack Trace Highlights
```text
panic: runtime error: invalid memory address or nil pointer dereference
goroutine 7 [running]:
[github.com/ethereum/go-ethereum/core/state.(*stateObject).setNonce](https://github.com/ethereum/go-ethereum/core/state.(*stateObject).setNonce)(...)
    /home/as/op-geth/core/state/state_object.go:570
[github.com/ethereum/go-ethereum/core/state.nonceChange.revert](https://github.com/ethereum/go-ethereum/core/state.nonceChange.revert)(...)
    /home/as/op-geth/core/state/journal.go:369
Proof of Concept
The issue can be reproduced by chaining CreateAccount, SelfDestruct, and RevertToSnapshot on the same address. See op_stack_revert_panic_test.go for the full implementation.

Impact
An attacker could potentially craft a transaction that, when processed by a sequencer or a validator node, causes the node to panic and shut down. In a network environment, this could lead to a halt in block production or synchronization issues.
