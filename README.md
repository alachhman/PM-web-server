# PokemasDB
A REST api for anyone to pull data for any (or all) Pokemon Masters EX characters. You can visit the homepage [here](http://pokemasdb.com) or explore the endpoints.

## Endpoints
### Trainer List
> https://pokemasdb.com/trainer <br />

Contains a list of every single trainer in the game, however **only the trainer's data.** 
Here's an example of what each entry in the list will look like.

```
{
  trainers: [
    {
      name: "Acerola",
      rarity: 5,
      pokemon: [
        "Palossand",
        "Mimikyu",
        "Mimikyu (Busted Form)"
      ],
      image: "https://pokemasdb.com/trainer/image/Acerola.png",
      data: "https://pokemasdb.com/trainer/Acerola"
    },
    ...109 more
  ]
}
```
___
### Specific Trainer
> https://pokemasdb.com/trainer/{TrainerName} <br />

Contains data on a specific trainer. To get in-depth data on any specific trainer just substitute `{TrainerName}` in the above url with the name of any trainer you want data for. **Keep in mind that the trainer names are case sensitive.** <br />

EX: https://pokemasdb.com/trainer/Red

This Endpoint contains the data for the trainer **as well as each of this trainer's pokemon, including seasonal variants, sygna suit pairings, and all evolutions.** Also, trainers that have spaces or other non-URL friendly characters in their name should have them changed to the HTML URL encoded version.

Examples of bad URLS:
* `https://pokemasdb.com/trainer/Lt. Surge` should be https://pokemasdb.com/trainer/Lt.%20Surge 
* `https://pokemasdb.com/trainer/SygnaSuitRed` should be https://pokemasdb.com/trainer/Red
___
### Trainer Image
> https://pokemasdb.com/trainer/image/{TrainerName}.png <br />

Contains a high quality, transparent background, PNG of each trainer. Substitute `{TrainerName}` in the above url with the name of any trainer you want an image for. The same rules for trainer name conventions in the `/trainer/{TrainerName}` endpoint also apply to this endpoint.

EX: http://pokemasdb.com/trainer/image/Red.png

**Pokemon Images Maybe Coming Soon**

## Example Usage
This api can be hit from any language you would like, however here's an example of how to grab the trainer list from the `/trainer` endpoint using node.js:
```javascript
const fetch = require("node-fetch");

const grabTrainer = async () => {
     let trainerUri = "https://pokemasdb.com/trainer";
     return await fetch(trainerUri)
        .then(response => response.json())
        .then(json => json.trainers)
};

grabTrainer().then(console.log);
```
## Wrappers
### Java
> [PM4J] (https://github.com/V-Play-Games/PM4J/tree/main)
Maven repository for use in Java Applications.

## Need Help?
For any additional questions, feature suggestions, or if you want to contribute: 

You can contact An🗡nee#0777 on Discord or join the [An-Two-Nee support server](https://discord.gg/cb8vna).
