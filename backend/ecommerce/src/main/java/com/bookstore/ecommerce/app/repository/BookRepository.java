package com.bookstore.ecommerce.app.repository;

import java.util.concurrent.CompletableFuture;

import com.bookstore.ecommerce.app.domain.data.Book;

public interface BookRepository {
  CompletableFuture<Book[]> find(String term);
}
