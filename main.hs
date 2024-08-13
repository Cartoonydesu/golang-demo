main = do
    print(isEven 3, isOdd 3)

isEven :: Integer -> Bool
isEven num = num `mod` 2==0

isOdd :: Integer -> Bool
isOdd num = num `mod` 2==1