# evesim
evesim is a toy program to experiment with the probability of a particular Mitochondrial DNA lineage dominating a population.

## What?
Mitochondrial DNA is only passed from a mother to her children.  If a mother has no children, or only a son, her mitochondrial DNA line dies with her.  There's a relatively recent discovery that all humans are descended from a single mitochondrial line, traced back to "Mitochondrial Eve".  Unlike her biblical namesake, it's generally accepted that "Mitochondrial Eve" did not live alone as the sole initial representative of the human race, but she is the only of her contemperaries (except her mother) whose produced an unbroken matrilineal line, from which all modern humans are descended.

## Why code?
This repository was born out of a conversation with my wife, who is reading Gene, by Siddhartha Mukherjee.  There is some debate between us about exactly how the proposition started, but I essentially started contemplating the following problem statement:  How likely is it that a given line of mitchondrial DNA will become dominant?  There are many discussions, graphics, and debates online, ranging from respected geneticists debating the history of early humans, to reddit comment threads.  I wanted to approach the problem from first principals, and writing some code felt like the right way to explore the problem space.  My starting assumptions are:
   1. There are N homo sapien females at T0.
   1. Their various genetic lines reproduce with equal fitness.
   1. Mitochondrial DNA is only inherited through female children.
   1. The likelihood of having female children is a random value, between 0 and the maximum number of female children (which turns out to be an important variable).

Given those conditions, how likely is a stable equilibrium of more than one mitochondrial line?  Contrary to my expectations, it seems the answer is, "not very likely".

## What I learned

The two critical variables are the number of initial females, and the maximum number of girls that can be produced in a given generation.  The mean "replacement rate" will be one half of the maximum, because of the assumption of a random number of children between 0 and max.  If the "replacement rate" of children is too low, over enough generations randomness wins out, and all of the genetic lines fail.  The necessary replacement rate to sustain an initial population of 10 females seems to be a maxGirls value of 3.0, or an average rate of 1.5 children for every mother.  Allowing for a larger starting population of females (and thus mitochondrial DNA lines) makes the populations less likely to die out, but reaffirms that one line will dominate on the way to very large population sizes.

Depending on the variables you choose, it's likely but possible that more than 1% of the population will have mitochondrial line which is different from the dominant line.  Unless you assume a very large starting population and/or a very high replacement rate, it's very unlikely that more than 10% of the population will have an alternate mitochondrial line.  The probability of having two stable lines, where the second-most dominant line captures more than 25% of the children, is typically less than a 1% chance.

## Example output

A sample run with the following values:
   1. initialPopulation := 10
   1. maxGirls := 3.1
~~~~
$ go run evesim.go
26% of the populations died out.
Averaged 76.564000 generations.
Last generation: 426.
Odds of a second lineage of > 1% of the population:  8.00 %
Odds of a second lineage of >10% of the population:  2.20 %
Odds of a second lineage of >25% of the population:  0.60 %
$
~~~~

## What could be improved
I strongly suspect that the likelihood of female children is not linear between 0 and maxGirls.  It's probably heavily weighted towards the bottom of that spectrum, probably a logarithmic scale.  It would be interesting to research real values for the number of children couples have at various points of societal development (I seem to recall it goes down in modern developed countries), and attempt to calculate a better formula for the number of children offset by the likelihood of having a girl.

Or I could just invesitage one of the several dozen simulations listed on the Coalescent theory Wikipedia page which seem to approach this problem from the opposite direction.
