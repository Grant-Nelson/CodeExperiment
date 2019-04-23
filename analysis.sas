DM 'LOG; CLEAR; OUT; CLEAR;';

DATA IN;
    INPUT order $ replicate $ index $ language $ algorithm $ seconds;

LINES;
{{ RESULTS GO HERE }}
;

PROC GLM DATA=IN PLOTS=ALL;
	CLASS language algorithm;
	MODEL seconds = language language(algorithm) / SS3;
	RANDOM language language(algorithm) / TEST;
    MEANS language language(algorithm);
TITLE '';

PROC GLM DATA=IN PLOTS=ALL;
	CLASS language algorithm;
	MODEL seconds = language|algorithm / SS3;
	RANDOM language algorithm / TEST;
    MEANS language|algorithm;
TITLE 'Factorieal Design';

RUN;
