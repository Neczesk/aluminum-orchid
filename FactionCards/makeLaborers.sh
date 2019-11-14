#!/usr/bin/env bash
echo "processing Laborers Faction Card"
rm -f LaborersFactionCard.tex
cp factionCardsTemplate.tex LaborersFactionCard.tex
sed -i 's/Faction Name/The Laborers/' LaborersFactionCard.tex
sed -i 's_\input{FactionDescriptions/loremipsum.tex}_\input{FactionDescriptions/Laborersdescription.tex}_' LaborersFactionCard.tex
sed -i 's_Ability 4:_Landscape Engineering:_' LaborersFactionCard.tex
sed -i 's_Ability 4\\input{FactionRules/loremipsum.tex}_\\input{FactionRules/landscapeengineering.tex}_'  LaborersFactionCard.tex
sed -i 's_Ability 5:_Clear Land:_' LaborersFactionCard.tex
sed -i 's_Ability 5\\input{FactionRules/loremipsum.tex}_\\input{FactionRules/clearland.tex}_'  LaborersFactionCard.tex
