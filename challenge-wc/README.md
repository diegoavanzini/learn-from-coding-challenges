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
vedi commit

vado ad implementare il metodo `CountBytes()` in modo da far passare il test
```
$ make test
go test .
ok      ccwc    1.006s
```
abbiamo il nostro primo test verde! 

Aggiungo il test che verifica, in caso di file errato, che il metodo ritorni un errore.

Infine é necessario chiamare il nostro eseguibile da linea di comando e far arrivare al nostro metodo i parametri che gli servono per fare quello che deve.

Per questo ci viene in aiuto il package `flag`.

In Go la funzione `main` é il punto di accesso del nostro tool ed é lei che riceve i parametri in input al tool.

Al fine di testare la validazione dei parametri in ingresso ho creato la funzione privata `validateInput` che riceve in ingresso i parametri passati e utilizzando il pacchetto `flag` definisce i parametri accettati e ne esegue il parsing e la validazione.

Il primo test verifica che senza parametri in ingresso la validazione ritorni un errore:

```golang
func TestWc_validateInputWithoutArguments_ShouldReturnError(t *testing.T) {
	// ARRANGE
	args := []string{}

	// ACT
	_, _, err := validateInput(args)
	defer resetFlags()

	// ASSERT
	assert.NotNil(t, err)
	assert.Equal(t, "count what?", err.Error())
}
```

In particolare con 

```golang
byteCount = flag.CommandLine.Bool("c", false, "count bytes in file usage: ccwc -c <file>")
```
dichiaro che uno dei parametri del mio tool sará un booleano che si chiamerá "c" ed di default sará `false`, mentre se presente sará `true`.

con 
```golang
flag.CommandLine.Parse(args)
```
vado a fare il parsing degli argomenti.

Gli argomenti in ingresso saranno quelli passati a line di comando `os.Args[1:]` togliendo il primo che é il nome del comando oppure quelli passati nel test.

Da notare la chiamata in defer della funzione `resetFlags()`, questa é necessaria in quando la dichiarazione di `CommandLine` é statica, questo fa si che quando lancio il secondo test che va a ridefinire `c` questo vada in errore ma noi non lo vogliamo. Per evitarlo lanciamo questo 
```golang
func resetFlags() {
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError) //flags are now reset
}
```

altri test che ho ritenuto utili sono quando il flag é presente ma non c´é il nome del file (`TestWc_validateInputWithCountFlagButNoFilepath_ShouldReturnError`) oppure c´é il nome del file ma é sbagliato (`TestWc_WhenInputFilePathIsWrong_ShouldReturnAnError`).


# step 2

in questo step dobbiamo calcolare il numero di righe in un file, utilizzando il file test.txt ci si aspetta (expected) 7145 righe.

Iniziamo come al solito dal test con il file di txt in dotazione e ritorno in prima battuta 0 facendo fallire il test.

Il test si aspetta il metodo `CountRows` associato alla struc `wc`, questo ritornerá il numero di linee. Per farlo deve anch'esso andare a leggere il file in input e vien da se che é necessario in questo caso entrare nel merito del contenuto del file per individuare gli a capo `[]byte{'\n'}`.

Sia che si tratti di leggere il numero di bytes o il numero di righe é necessaria la lettura del file. Per leggere il file si usa il descrittore ritornato da `os.Open(w.filepath)` e utilizzando il metodo `Read` andiamo a leggere e mettere il contenuto in un array di byte.

Centralizzo il calcolo in un metodo che si occupa di leggere il file e instanaziare i contatori che ci interessano.

```golang
func (w *wc) readFile(inputFile string) error {
	r, err := os.Open(w.filepath)
	if err != nil {
		return err
	}
	buf := make([]byte, 32*1024)
	for {
		c, err := r.Read(buf)
		w.numberOfLine += bytes.Count(buf[:c], []byte{'\n'})
		w.numberOfBytes += c
		switch {
		case err == io.EOF:
			return nil
		case err != nil:
			return err
		}
	}
}
```
ora é necessario introdurre il flag a linea di comando che dice di calcolare il numero di linee quindi altro test rosso che andiamo a mettere a posto.
E se inserisco piú flag ritorno  il conteggio di linee e bytes.