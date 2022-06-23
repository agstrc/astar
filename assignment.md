# Assignment

The following is a translation of the main parts of the original assignment.  

*"After killing the king of Hyrule, the mage Agahnim is keeping the Zelda princess as a
prisioner e pretends to break the seal which keep the evil Ganon imprisioned in the Dark
World. Link is the only warrior able to win agains the mage Agahnim, save Zelda and bring
piece to the kingdom of Hyrule. The only weapon strong enough to defeat the mage Agahnim
is the legendary Master Sword, which is attached to a pedestal in Lost Woods. To prove
that he is worthy of wielding the Master Sword, Link must find and retrieve the three
Pendants of Virtue: courage, power and wisdom. The three pendants are spread through the
kingdom of Hyrule, inside the dangerous dungeons."*

## Goal

The assignment consists in implementing an agent capable of autonomous traversal through
the kingdom of Hyrule, exploring the dangerous dungeons and collecting the Pendants of
Virtue. To do so, you must use the A* algorithm. The agent must be capable of
automatically calculate the best route to collect the three Pendants of Virtue and go to
the Lost Woods, where the Master Sword is located.  

The kingdom of Hyrule consists of 5 terrain types: grass, water, mountains, sand and
forests. The costs to traverse each of these terrains are:

- Grass -- cost: 10
- Sand -- cost: 20
- Forest -- cost: 100
- Mountains -- cost: 150
- Water -- cost: 180

Within the dungeons, traversal is only possible in the lighter areas. The cost to walk
through this kind of terrain is 10.  

Additional information:

- Before reaching the entrance of the Lost Woods, the agent must start his journey at a
starting point and return to said starting point with all pendants collected.
- The agent may not move in a diagonal orientation.
