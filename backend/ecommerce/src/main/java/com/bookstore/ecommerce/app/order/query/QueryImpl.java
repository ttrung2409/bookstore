package com.bookstore.ecommerce.app.order.query;

import java.util.Arrays;
import java.util.concurrent.CompletableFuture;
import com.bookstore.ecommerce.app.order.query.model.Book;
import com.bookstore.ecommerce.app.order.query.model.Order;
import com.bookstore.ecommerce.app.repository.BookRepository;
import com.ea.async.Async;

import org.springframework.stereotype.Component;


@Component
public class QueryImpl implements Query {
  private final BookRepository bookRepository;

  public QueryImpl(BookRepository bookRepository) {
    this.bookRepository = bookRepository;
  }

  @Override
  public CompletableFuture<Book[]> findBooks(String term) {
    var books = Async.await(this.bookRepository.find(term));

    return CompletableFuture.completedFuture(
      Arrays.stream(books)
        .map(book -> Book.fromDataObject(book))
        .toArray(Book[]::new));
  }

  @Override
  public Order getOrderDetails(String orderId) {
    return null;
  }
}
