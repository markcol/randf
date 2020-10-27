# randf

A better floating-point random number generator.

This package generates uniformly-distributed pseudo-random 32-bit
floating-point values in the range [0, 1], assuming that we are given an
algorithm for generating pseudo-random bits. These bits should have
probability 50% of being in each of two states, and be statistically
independent.

The approach is to choose floating-point values in the range such that the
probability that a given value is chosen is proportional to the distance
between it and its two neighbors. The algorithm works in two steps, first
choosing the exponent range of the value, then choosing the mantissa.

From the paper <http://allendowney.com/research/rand/downey07randfloat.pdf>.
