# Daily Coding Problem: Problem #617 [Medium]

This problem was asked by Facebook.

Given a number in Roman numeral format, convert it to decimal.

The values of Roman numerals are as follows:

    {
        'M': 1000,
        'D': 500,
        'C': 100,
        'L': 50,
        'X': 10,
        'V': 5,
        'I': 1
    }

In addition, note that the Roman numeral system uses subtractive notation for numbers such as IV and XL.

For the input XIV, for instance, you should return 14.

## Build and run

This is a command line program in Go.

    $ go build .
    $ ./romannumerals XIV
    XIV -> 14

## Analysis

I implemented this as a state machine of sorts,
which keeps track of a sum (which will ultimate be the answer),
a run (the sum of some digits) which hasn't yet been added  or subtracted to the sum,
and the last digit's value.

Based on the current and the last digit's values,
the machine:

* adds to the run, if the current digit's value equals the last digit's value
* adds run to sum, sets run to current digit's value, if the current digit value is less than last digit's
* adds current digit - run to sum, sets run to 0, if the current digit has greater value than last digit

I don't really know if this is the best way.
It lets through some dubious roman numerals, like IIIIX.

This might constitute a decent interview question,
if you're deeply concerned about your candidate knowing the mechanics
of the programming language of the job.
Getting it correct will require a lot of finicky tests, looping,
and if-then-else constructs.
Since the input is a string, some small amount of string processing will appear, as well.

The drawback I can see is that Roman Numerals have no great definition,
so corner cases abound.
The candidate may get bogged down thinking of odd corner cases
rather than writing some code.
If you're a candidate,
you should beware of that trap, too.

The candidate could score points by noting drawbacks to whatever algorithm,
like giving an answer for IIIIX.
Maybe note things like that as test cases.

I'm not sure how I'd rate this one,
so "Medium" seems like an appropriate rating.
