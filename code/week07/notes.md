# Week 7 State Machines

data State Machine constructor takes 4 parameters
[haddock documentation](http://localhost:5000/plutus-contract/html/Plutus-Contract-StateMachine)
- `smTransition` transistion function to change a state into another state
Input arguments:
  - state
    - stateData (Datum)
    - stateValue (Value)
  - Redeemer
  - If a Transition is not allowed we can pass Nothing as the third Parameter. if it is Allowed we pass a tuple
    - TxConstraints (Additional Constraints for the transition)
    - New State (new Datum and Value)


- `smfinal` tells wheter the state is Final (no way out) or not

- `smCheck` it just gets the datum redeemer and the contexts (additional checks that cannot be represented by TxConstraints)
- `smThreadToken` can or cannot be there (We mint an NFT and because this can only exist once)