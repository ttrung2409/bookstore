package com.bookstore.ecommerce.repository.elastic;

import java.util.ArrayList;
import java.util.List;
import java.util.concurrent.CompletableFuture;

import com.bookstore.ecommerce.app.domain.data.Book;
import com.bookstore.ecommerce.app.repository.BookRepository;

import org.springframework.stereotype.Component;

@Component
public class ElasticBookRepository implements BookRepository {
  @Override
  public CompletableFuture<List<Book>> find(String term) throws Exception {
    return CompletableFuture.completedFuture(new ArrayList<Book>());
  }
}
