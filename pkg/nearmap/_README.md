## Nearmap Ref Docs
-  [Features](https://docs.nearmap.com/display/ND/Nearmap+AI+Feature+API)
- [Api](https://api.nearmap.com/ai/features/v4)

## Nearmap Roof materials

- shingle (roof material)
- tile (roof material)
- metal (roof material)

## Meeting  Mar 22, 2003

- [x] Pull survey before spring i.e. YYYY-04-01 (to avoid tree overhand issue)
- [x] If pitch degrees of primary roof is nil, pull previous survey(Max 5 attempts)
- [x] If material is shingle and no pitch use 5/12 as default pitch

## Jason provide following formula to calculate Total sq. ft from Flat sq. ft:

### Roofing Pitch from degree

``` text
    Formulas - Roofing Pitch, Multiplier, Total SqFt
    Roofing Pitch is X/Y where Y=12 (X/12)
    Solve for X where P = 30 degrees
    Formula
        X=tan(P)*12
        X= 6.93
    Variables
        P = Roof Pitch (degrees)
        X = Rise (2 decimal places)
        Y = Run = 12
```

### Actual area with roofing pitch

``` text
    Total SqFt is Flat SqFt * Multiplier
    Solve for T where P = 30 degrees and F = 3100 ft^2
    Solve for M
    Formulas
        T=(F/Y)*(√(x^2+y^2))
        T = 3580 ft^2
        M = 1 / cos(P)
        M = 1.155
    Variables
        F = Flat SqFt
        T = Total SqFt
        M = Multiplier (3 decimal places)
        P = Roof Pitch (degrees)
    Details
        F=W*Y
        T=W*H
        T=F*M
        H=√(x^2+y^2)
        W = 3100/Y
        T=(3100/12)*√6.93^2+144
        T = 3580 ft^2
```


