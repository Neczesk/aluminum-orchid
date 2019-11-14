#!/usr/bin/env bash
echo "processing White Robes Faction Card"
rm -f WhiterobesFactionCard.tex
cp factionCardsTemplate.tex WhiterobesFactionCard.tex
sed -i 's/Faction Name/The White Robes/' WhiterobesFactionCard.tex
sed -i 's_\input{FactionDescriptions/loremipsum.tex}_\input{FactionDescriptions/whiterobesdescription.tex}_' WhiterobesFactionCard.tex
sed -i 's_Ability 4:_Research Technology:_' WhiterobesFactionCard.tex
sed -i 's_Ability 4\\input{FactionRules/loremipsum.tex}_\\input{FactionRules/researchtechnology.tex}_'  WhiterobesFactionCard.tex
sed -i 's_Ability 5:_Write Chronicle:_' WhiterobesFactionCard.tex
sed -i 's_Ability 5\\input{FactionRules/loremipsum.tex}_\\input{FactionRules/writechronicle.tex}_'  WhiterobesFactionCard.tex
