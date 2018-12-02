Word List Problem
Identify all combinations where one words can be composed of two or more words of the same list, and print them.

Problem
You are given a (potentially large) List of words. Some of these are compounds, where all parts are also in the List. Example List:

[rockstar, rock, star, rocks, tar, star, rockstars, super, highway, high, way, superhighway]
The task is to identify all combinations where one word is a composite of two or more words from the same list and print or return them.

Example:

rockstar -> rock + star
superhighway -> super + highway
superhighway -> super + high + way
If returning, the output would be a list of lists, e.g.

[[rock, star], [super, highway], [super, high, way],...]
