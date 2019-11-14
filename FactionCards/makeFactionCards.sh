#!/usr/bin/env bash

LOG=log.log
rm log.log
echo "making faction cards" >> $LOG

./makeGuilds.sh >> $LOG
./makeWhiteRobes.sh >> $LOG
./makeLaborers.sh >> $LOG
./makeBlackRobes.sh >> $LOG
pdflatex BlackRobesFactionCard.tex >> $LOG
pdflatex WhiteRobesFactionCard.tex >> $LOG
pdflatex LaborersFactionCard.tex >> $LOG
pdflatex GuildsFactionCard.tex >> $LOG
