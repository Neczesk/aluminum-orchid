#!/usr/bin/env bash
echo "processing Guilds Faction Card"
rm -f GuildsFactionCard.tex
cp factionCardsTemplate.tex GuildsFactionCard.tex
sed -i 's/Faction Name/The Guilds/' GuildsFactionCard.tex
sed -i 's_\input{FactionDescriptions/loremipsum.tex}_\input{FactionDescriptions/guildsdescription.tex}_' GuildsFactionCard.tex
sed -i 's_Ability 4:_Craft Gear:_' GuildsFactionCard.tex
sed -i 's_Ability 4\\input{FactionRules/loremipsum.tex}_\\input{FactionRules/craftgear.tex}_'  GuildsFactionCard.tex
sed -i 's_Ability 5:_Train Master:_' GuildsFactionCard.tex
sed -i 's_Ability 5\\input{FactionRules/loremipsum.tex}_\\input{FactionRules/trainmaster.tex}_'  GuildsFactionCard.tex
