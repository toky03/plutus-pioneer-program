# Notes Week 5 Native Tokens
For Plutus use directly the Tag from Week 6
`8a20664f00d8f396920385947903761a9a897fe0`

`import Plutus.V1.Ledger.Ada`

`:set -XOverloadedStrings` to enter byteStrings as literal Strings

### Helperfuntions to get a Token value
`lovelaceValueOf`
Gives a Map from tokenmap to integers


The Value Class is an instance of Monoid thus they are things that can be combined.
The operator to combine two Values is `<>`

## Create Values
`import Plutus.V1.Ledger.Value`
```haskell
:t singleton
singleton :: CurrencySymbol -> TokenName -> Integer -> Value
```
CurrencySymbol must represent a hexadecimal value
```haskell
:t valueOf
valueOf :: Value -> CurrencySymbol -> TokenName -> Integer
```
```haskell
:t flattenValue
flattenValue :: Value -> [(CurrencySymbol, TokenName, Integer)]
```

## Why do we need a currency symbol and a token name?
Because of the minting policies
fees depend on the size in bytes
A currency Symbol is actually the Hash of a script (therefore it must be represented by a hexadecimal value)
The script is called the minting policy and the corresponding script must also be in the transaction.

ADA has no Currency Symbol, which means there is no script which can be executet. Therefore the total ammount of ADA is fixed

For each Native 
> ADA can never be minted or burn. Only Custom Tokens can be burnt or minted
>Tipp: To prevent others from creating a Token with the same Currency Symbol, one can use the Wallet PubKeyHash as Currency Symbol

# NFT Non Fungible Tokens
## Previous creation
Minting is only allowed before a Deadline
with a Blockchain explorer it is checked that only one Token is existing

## creation with Plutus
We use a UTXO to guarantee that
We need to make sure, we consume the specified UTXO
```haskell
any (\i -> txInInfoOutRef i == oref) $ txInfoInputs $ scriptContextTxInfo ctx
```
> any is a standart Haskell function which checks if any of the elements satisfies a certain condition
