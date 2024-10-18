# Building wc
# step 0

Iniziando dalla  challenge https://codingchallenges.fyi/challenges/challenge-wc del contatore di parole ho implementato lo step zero.

Creato il main vuoto, un test vuoto, un file "wc.go" che conterrá tutta la logica, inizializzato il modulo "wc" e la cartella "target" con il compilato che attualemnte non fa niente.

# step 1

in questo step il tool "ccwc" deve restituire il numero di bytes in un file, utilizzando il file test.txt ci si aspetta (expected) 342190 bytes. 

il package `os` mette a disposizione il metodo `ReadFile` che restituisce un array di byte e va da se che contarli é semplice come contare la `len` di un array.

```golang
content, err := os.ReadFile(inputFile)
if err != nil {
    return nil, err
}
```
la prima cosa che ho creato cercando un approccio di tipo [TDD](https://it.wikipedia.org/wiki/Test_driven_development) creando il test `TestWc_WhenInputFile_ShouldReturnTheExpectedNumberOfBytes` mettendoci dentro quello che mi aspetto e cercando di farlo compilare. In particolare ho creato una struttura `wc` con il suo costruttore `NewWc` che riceve il path del file di cui deve contare i bytes ed espone il metodo `CountBytes()`. In prima battuta il metodo ritornerá 0 bytes mentre nell'assert ce ne aspettiamo 342190 se utilizziamo il file `test.txt`... il test quindi fallirá.
```
--- FAIL: TestWc_WhenInputFile_ShouldReturnTheExpectedNumberOfBytes (0.00s)
    c:\GoCode\src\github.com\diegoavanzini\learnfromcodechallenges\challenge-wc\wc_test.go:17: 
        	Error Trace:	c:/GoCode/src/github.com/diegoavanzini/learnfromcodechallenges/challenge-wc/wc_test.go:17
        	Error:      	Not equal: 
        	            	expected: 342190
        	            	actual  : 0
        	Test:       	TestWc_WhenInputFile_ShouldReturnTheExpectedNumberOfBytes
FAIL
FAIL	ccwc	1.183s
FAIL
```


un pó piú complicato é gestire i parametri in ingesso al nostro tool