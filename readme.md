# Ilmalaiva
**Ilmalaiva** (fin. "Airship") is small game made for gamejam hold on
[Lyof's Code Constellation server](https://discord.gg/khZnX8TRHU) with theme
"in the sky".

The concept behind Ilmalaiva was to make airship-based peaceful sim.. but
the time for gamejam was rough, so the game took more extreme direction.

## Gameplay
You are on airship. I have no idea why, but you have one mission: to keep it
alive. The stove is burning energy fast, so you need to head to nearest generator
as fast as possible!

|               Construction              |                   Description                   |
|:---------------------------------------:|:------------------------------------------------|
|   ![](assets/tiles/special_stove.png)   | **Stove** - requires fuel to keep airship alive |
| ![](assets/tiles/special_generator.png) | **Generator** - produces energy for stove       |

**Movement**  
Move yourself with `WASD` or `arrow keys`.<br>
Using `space` on generator allows you to get energy. Use `space` on stove to
resupply it!<br>

**Boost**  
Every stove resupplying count into consecutive boost bonus. Once you gather two
or more, you can use `Q` to get faster speed or slowing stove down.  
The bigger the count, the more powerful the bonus, and with count big enough
you can even make small combos!

## Customising?
You can create your own map by writing `.ilmp` file and putting it into `maps` folder!
Then, simply write its name into `ilmalaiva.yaml` file under `map` keyword.  
Creating map file is easy: simply fill document with respective symbols, where
every row and line matches respective coordinates. You can see list of symbols
[here](ilmp_format.md). Remember that one `P` symbol is required (player spawnpoint)
and having stoves/generators is needed for the game to be possible to play.

You can also customise background by adding new image into `skies` folder and then
using its name in `ilmalaiva.yaml` file under `background` keyword.  
Try to get image with similar size to `default.png` to fit game resolution.
