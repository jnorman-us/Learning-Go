# Learning-Go
I decided to learn Golang by creating a solver for the Cracker Barrel Peg Game

## Running
From the Root Directory, run the following command:

    go run main.go
  
You will be prompted to enter a 15 letter string of the characters `0` and `1`; 0 means empty, 1 means there is a peg. 
The 15 character stream maps to the Cracker Barrel Triangle like this:

            [ 0]
          [ 1] [ 2]
        [ 3] [ 4] [ 5]
      [ 6] [ 7] [ 8] [ 9]
    [10] [11] [12] [13] [14]
    
Enter the 15 characters of the board you would like solved:

    josephnormandev/Learning-Go~$ go run main.go
    Enter a Board: (15 characters long)
    ...............
    111100110111111

Then press enter!

## Results
The results for this board are pretty interesting; it leaves two pegs on opposite ends of the board. Weird!

    Results........
    Round: 0
        *
       * *
      * - -
     * * - *
    * * * * *

    Round: 1
        *
       * *
      * * -
     * - - *
    * - * * *

    Round: 2
        -
       * -
      * * *
     * - - *
    * - * * *

    Round: 3
        *
       - -
      - * *
     * - - *
    * - * * *

    Round: 4
        *
       - *
      - * -
     * - - -
    * - * * *

    Round: 5
        *
       - *
      * * -
     - - - -
    - - * * *

    Round: 6
        *
       - -
      * - -
     - * - -
    - - * * *

    Round: 7
        *
       - -
      * - -
     - * - -
    - * - - *

    Round: 8
        *
       - -
      - - -
     - - - -
    - * * - *

    Round: 9
        *
       - -
      - - -
     - - - -
    - - - * *

    Round: 10
        *
       - -
      - - -
     - - - -
    - - * - -
