#!/usr/bin/env bash
echo "processing Black Robes Faction Card"
rm -f BlackRobesFactionCard.tex
cp factionCardsTemplate.tex BlackRobesFactionCard.tex
sed -i 's/Faction Name/The Black Robes/' BlackRobesFactionCard.tex
sed -i 's_\input{FactionDescriptions/loremipsum.tex}_\input{FactionDescriptions/blackrobesdescription.tex}_' BlackRobesFactionCard.tex
sed -i 's_Ability 4:_Cast Ritual:_' BlackRobesFactionCard.tex
sed -i 's_Ability 4\\input{FactionRules/loremipsum.tex}_\\input{FactionRules/castritual.tex}_'  BlackRobesFactionCard.tex
sed -i 's_Ability 5:_Learn Magic:_' BlackRobesFactionCard.tex
sed -i 's_Ability 5\\input{FactionRules/loremipsum.tex}_\\input{FactionRules/learnmagic.tex}_'  BlackRobesFactionCard.tex
