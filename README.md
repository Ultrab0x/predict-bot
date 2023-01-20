# predict-bot

This script illustrates how a prediction bot could work on Uniswap. It connects to a local Ethereum node on port 8545 and creates a new Uniswap exchange instance using the exchange address. It then defines a prediction model, which is a what-if structure with a single PredictPrice function that takes the current price and returns a predicted price. In a loop, it gets the current token price, gets the predicted model price, and compares the two. If the predicted price is higher than the current price, it will supplement, if it is lower, it will remove
