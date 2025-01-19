# Build Your Own JSON Parser
https://codingchallenges.fyi/challenges/challenge-json-parser/

## Step 1

In this step we have to check the brackets in the input json. This remember to me an exercises done in Leetcode ([here my solution](https://github.com/diegoavanzini/go-grind75/tree/main/valid-brackets) ).

We have to parse the input and check for open bracket with a map which presents the link between open e correspondent closed bracket. Then for every open bracket we put in a stack the closed bracket I expect to find next.

## Step 2

In the second step we have to parse the string between brackets and in this simplest case we have only a key and a value separated by colon. We use ``strings.Split`` method and the related tests are green. Then we start some refactoring. 

## Step 3

In this step we use the method ``parseSingleKeyValue`` but before we have to split by comma the string between the brackets to get the single key value pair

## Step 4

Il quarto step ci mette in crisi perché nel parsing che facciamo del json andiamo a leggere tutto il contenuto mentre dovremmo pare un parsing parziale per leggere i singoli elementi che lo compongono... cosí da individuare e gestire la lettura degli stessi. Per fare questo dobbiamo rivedere il modo di leggere senza rompere i test che abbiamo giá scritto (cioé senza regressioni).
Quello che mi viene in mente di fare é di leggere il contenuto tra la prima parentesi e l'ultima, controllare se corrispondono, leggere le chiavi e i relativi valori ed eseguire ricorsivamente la lettura del contenuto quando trovo una parentesi graffa aperta (oggetto).   