# Go http averages kata

The goal of this kata is to improve your skills when it comes to testing
http services in golang. The application to be build will calculate the 
daily percentage-wise change in stock prices for a particular stock and forward
this result to another service. You may implement the receiving endpoint
in whichever you way you see fit. However, it should receive JSON in the following
format:

```json
{
"msf": [
["2020-02-20", 120],
["2020-02-21", 100],
["2020-02-22", 110],
["2020-02-23", 130]
],
"goog": [
["2020-02-20", 220],
["2020-02-21", 250],
["2020-02-22", 210],
["2020-02-23", 180]
]
}
```

This means the stock msf was worth 120$ on 2020-02-20 and 100$ on 2020-02-21 etc. Note that
there can be more than one stock item (in this example msf and goog).

The endpoint for the other service is 
POST http://foo.invalid/changes/{Symbol} where {Symbol} is a ticker symbol 
(for example MSF for microsoft). It must receive a list of changes in percentages per day.
Example: 
POST http://foo.invalid/changes/MSF
Body:
```json
[
["2020-02-21", -16.67],
["2020-02-22", 10],
["2020-02-23", 18.18]
]
```

The first value means that from the 2020-02-20 to 2020-02-21 the stock went down by -16.67
percent. The second value means that the stock went up by 10 percent from 2020-02-21 to 2020-02-22.
Note that 2020-02-20 is missing. This is because there is no value before that so we cannot
calculate a change in percentage. 
