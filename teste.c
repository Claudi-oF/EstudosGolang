#ifndef LISTA_H_INCLUDED
#define LISTA_H_INCLUDED

#define MAX 100

struct aluno{
    int matricula; 
    char nome[30];
    float n1, n2, n3;
};

typedef struct lista Lista;

Lista * criarLista(); 
void liberarLista(Lista * li); 

int buscaListaPosicao(Lista * li, int pos, struct aluno * al);
int buscaListaMatricula(Lista * li, int pos, struct aluno * al);

int insereListaInicio(Lista * li, struct aluno * al);
int insereListaOrdenado(Lista * li, struct aluno * al);
int insereListaFinal(Lista * li, struct aluno * al);

int removeListaInicio(Lista * li);
int listaVazia(Lista * li);
int listaCheia(Lista * li);