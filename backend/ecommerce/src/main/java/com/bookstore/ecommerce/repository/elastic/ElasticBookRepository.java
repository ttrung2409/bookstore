package com.bookstore.ecommerce.elastic;

import java.util.concurrent.CompletableFuture;

import com.bookstore.ecommerce.app.domain.data.Book;
import com.bookstore.ecommerce.app.repository.BookRepository;

import org.springframework.stereotype.Component;

@Component
public class ElasticBookRepository implements BookRepository {
  @Override
  public CompletableFuture<Book[]> find(String term) {
    return null;
  }
}
