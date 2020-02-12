Skills & Scripts
================

In this folder there are two types of files:

- character skill manifests
- skill scripts

## Skill Manifest

These are YAML files that describe all the skills for each class, including the
basic attack they all start with.

The manifest is laid out like so:

```
<uuid>:
  name: <string>
  type: <type>
  description: 
    1: <string>
    2: <string>
    3: <string>
  cooldown: <int>
  cost: <int>
  levelArgs:
    1: []
    2: []
    3: []
  scripts:
    onActivate: <string>
    onAttack: <string>
    onMove: <string>
    onCooldown: <string>
    onRoundOver: <string>
    onLevelUp: <string>
```

* **uuid** is the ID of the skill
* **name** is what will be shown to the players
* **type** can be one of `Basic`, `Action`, or `Passive`
* **description** can either be a map with the keys 1,2 & 3 or a single string. 
  * if a single string, that will be the description of the skill for all three levels
  * if a map, that will be the description for each level ( useful for skills
    where the description should show how it improves next level )
* **cooldown** an integer that specifies how many rounds before the skill is
  active again. Same as description, can either be a single integer or a map --
  a map specifies the cooldown for each level.
* **cost** the Focus cost of using the skill
* **levelArgs** an array of things to be passed into the skill script when it is
  run
* **scripts** scripts that are registered with the character that will run when
  the appropriate trigger happens:
  * `onActivate` is when the user chooses the skill on their sheet 
  * `onAttack` is for passive skills that boost or modify the basic attack
  * `onDefense` is for skills that activate when the character is attacked
  * `onMove` is for skills that need to do something when the player moves ( ie,
    stopping an effect on a monster when the player moves away )
  * `onCooldown` is for the end of a round when cooldowns would normally be
    calculated. Use this for skills that have special requirements for when they
    reactivate ( such as skills that only reactivate when the target is dead, or
    at the end of a round )
  * `onRoundOver` is for skills that need to do something special at the end of
    a round
  * `onLevelUp` is for skills that need to add or change modifiers when the
    character levels up


## Scripts

Based on when a script is called, the arguments sent will change:

* `onActivate`
  * an array of UUIDs that represent who the skill is targeting

* `onAttack`
  * UUID of what is being attacked

* `onDefense`
  * UUID of what is attacking

* `onMove`
  * the integer of the threat zone the actor is leaving
  * the integer of the threat zone the actor is entering

* `onCooldown`

* `onRoundOver`
  * the round number

* `onLevelUp`
  * the new level of the actor
