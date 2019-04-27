DM 'LOG; CLEAR; OUT; CLEAR;';

DATA IN;
    INPUT order $ replicate $ index $ language $ algorithm $ seconds;

LINES;
{{ RESULTS GO HERE }}
;

PROC GLM DATA=IN PLOTS=ALL;
	CLASS language algorithm;
	MODEL seconds = language|algorithm / SS3;
	RANDOM language|algorithm / TEST;
    MEANS language|algorithm;
TITLE 'Language and Algorithm - Factorial Design';
RUN;

PROC GLMPOWER DATA=IN;
	 CLASS language algorithm;
     MODEL seconds = language|algorithm;
     POWER
    	 stddev = {{ Set Standard Deviation }}
    	 alpha = 0.05
    	 Ntotal = .
    	 Power = .88 .90 .92;   	 
TITLE 'Power for language|algorithm';
RUN;
