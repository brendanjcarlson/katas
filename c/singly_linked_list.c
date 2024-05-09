#include <stdio.h>
#include <stdlib.h>

typedef struct Node {
    int data;
    struct Node* next;
} Node;

void insert_end(Node** head, int value) {
    Node* new_node = malloc(sizeof(Node));
    if (new_node == NULL) {
        exit(1);
    }

    new_node->next = NULL;
    new_node->data = value;

    if (*head == NULL) {
        *head = new_node;
        return;
    }

    Node* curr = *head;
    while (curr->next != NULL) {
        curr = curr->next;
    }
    curr->next = new_node;
}

void insert_beginning(Node** head, int value) {
    Node* new_node = malloc(sizeof(Node));
    if (new_node == NULL) {
        exit(1);
    }
    new_node->data=value;
    new_node->next = *head;
    *head = new_node;
}

void insert_after_node(Node* node, int value) {
    Node* new_node = malloc(sizeof(Node));
    if (new_node == NULL) {
        exit(1);
    }
    new_node->data = value;
    new_node->next = node->next;
    node->next = new_node;
}

void insert_in_order(Node** head, int value) {
    if (*head == NULL || (*head)->data >= value) {
        insert_beginning(head, value);
        return;
    }

    Node* curr = *head;
    while (curr->next != NULL) {
        if (curr->next->data >= value) {
            break;
        }
        curr = curr->next;
    }

    insert_after_node(curr, value);
}

void remove_first_occurrance(Node **head, int value) {
    if (*head == NULL) {
        return;
    }

    if ((*head)->data == value) {
        Node* to_remove = *head;
        *head = (*head)->next;
        free(to_remove);
        return;
    }

    for (Node* curr = *head; curr->next != NULL; curr = curr->next) {
        if (curr->next->data == value) {
            Node* to_remove = curr->next;
            curr->next = curr->next->next;
            free(to_remove);
        }
    }
}

void reverse_linked_list(Node **head) {
    Node* prev = NULL;
    Node* curr = *head;

    while (curr != NULL) {
        Node* next = curr->next;
        curr->next = prev;
        prev = curr;
        curr = next;
    }
    *head = prev;
}

void display_linked_list(Node **head) {
    for (Node* curr = *head; curr != NULL; curr = curr->next) {
        printf("%d -> ", curr->data);
    }
    printf("\n");
}

void deallocate_linked_list(Node** head) {
    Node* curr = *head;
    while (curr != NULL) {
        Node* prev = curr;
        curr = curr->next;
        free(prev);
    }
    *head = NULL;
}

int main(int argc, char* argv[]) {
    Node* head = NULL;

    insert_end(&head, 6);
    insert_end(&head, 8);
    insert_end(&head, 10);

    insert_beginning(&head, 4);
    insert_beginning(&head, 0);

    insert_after_node(head, 2);

    insert_in_order(&head, 1);
    insert_in_order(&head, 7);
    insert_in_order(&head, 11);

    remove_first_occurrance(&head, 8);
    remove_first_occurrance(&head, 0);

    display_linked_list(&head);

    reverse_linked_list(&head);

    display_linked_list(&head);

    deallocate_linked_list(&head);
    return 0;
}
