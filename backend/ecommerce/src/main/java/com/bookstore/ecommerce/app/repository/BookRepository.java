package com.bookstore.ecommerce.app.repository;

import java.util.List;
import java.util.concurrent.CompletableFuture;
import com.bookstore.ecommerce.app.domain.data.Book;

public interface BookRepository {
  CompletableFuture<List<Book>> find(String term) throws Exception;
}
