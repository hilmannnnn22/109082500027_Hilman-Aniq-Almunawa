#ifndef QUEUE_H
#define QUEUE_H

#include <iostream>
using namespace std;

const int MAKSIMAL = 5;

typedef int infotype;

struct queue {
    infotype info[MAKSIMAL];
    int head;
    int tail;
};

bool isEmptyQueue(queue Q);
bool isFullQueue(queue Q);
void CreateQueue(queue &Q);
void enqueue(queue &Q, infotype x);
infotype dequeue(queue &Q);
void printInfo(queue Q);

#endif