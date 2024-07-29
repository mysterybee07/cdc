#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <stdbool.h>

#define MAX_SYMBOLS 256
#define MAX_PRODUCTIONS 128
#define MAX_PRODUCTION_LENGTH 64

typedef struct {
    char nonTerminal;
    char productions[MAX_PRODUCTIONS][MAX_PRODUCTION_LENGTH];
    int productionCount;
} Grammar;

typedef struct {
    bool symbols[MAX_SYMBOLS];
} FirstSet;

Grammar grammar[MAX_SYMBOLS];
FirstSet firstSets[MAX_SYMBOLS];
bool computedFirst[MAX_SYMBOLS];
int grammarSize = 0;

void initializeFirstSets() {
    for (int i = 0; i < MAX_SYMBOLS; i++) {
        for (int j = 0; j < MAX_SYMBOLS; j++) {
            firstSets[i].symbols[j] = false;
        }
        computedFirst[i] = false;
    }
}

int findNonTerminalIndex(char nonTerminal) {
    for (int i = 0; i < grammarSize; i++) {
        if (grammar[i].nonTerminal == nonTerminal) {
            return i;
        }
    }
    return -1;
}

void addFirstSet(int nonTerminal, char symbol) {
    firstSets[nonTerminal].symbols[(unsigned char)symbol] = true;
}

bool hasEpsilon(int nonTerminal) {
    return firstSets[nonTerminal].symbols[(unsigned char)'ε'];
}

void computeFirstSet(int nonTerminal);

void addFirstSets(int from, int to) {
    for (int i = 0; i < MAX_SYMBOLS; i++) {
        if (firstSets[from].symbols[i] && i != 'ε') {
            addFirstSet(to, i);
        }
    }
}

void computeFirstSet(int nonTerminal) {
    if (computedFirst[nonTerminal]) {
        return;
    }
    computedFirst[nonTerminal] = true;

    for (int i = 0; i < grammar[nonTerminal].productionCount; i++) {
        char *production = grammar[nonTerminal].productions[i];
        bool epsilonInProduction = true;

        for (int j = 0; j < strlen(production); j++) {
            char symbol = production[j];
            int symbolIndex = findNonTerminalIndex(symbol);

            if (symbolIndex == -1) {
                addFirstSet(nonTerminal, symbol);
                epsilonInProduction = false;
                break;
            } else {
                computeFirstSet(symbolIndex);
                addFirstSets(symbolIndex, nonTerminal);

                if (!hasEpsilon(symbolIndex)) {
                    epsilonInProduction = false;
                    break;
                }
            }
        }

        if (epsilonInProduction) {
            addFirstSet(nonTerminal, 'ε');
        }
    }
}

void printFirstSet(int nonTerminal) {
    printf("FIRST(%c) = { ", grammar[nonTerminal].nonTerminal);
    for (int i = 0; i < MAX_SYMBOLS; i++) {
        if (firstSets[nonTerminal].symbols[i]) {
            printf("%c ", i);
        }
    }
    printf("}\n");
}

int main() {
    grammar[0].nonTerminal = 'S';
    strcpy(grammar[0].productions[0], "AB");
    grammar[0].productionCount = 1;

    grammar[1].nonTerminal = 'A';
    strcpy(grammar[1].productions[0], "aA");
    strcpy(grammar[1].productions[1], "ε");
    grammar[1].productionCount = 2;

    grammar[2].nonTerminal = 'B';
    strcpy(grammar[2].productions[0], "bB");
    strcpy(grammar[2].productions[1], "ε");
    grammar[2].productionCount = 2;

    grammarSize = 3;

    initializeFirstSets();

    for (int i = 0; i < grammarSize; i++) {
        computeFirstSet(i);
    }

    for (int i = 0; i < grammarSize; i++) {
        printFirstSet(i);
    }

    return 0;
}
