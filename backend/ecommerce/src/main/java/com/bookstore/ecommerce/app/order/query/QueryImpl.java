package com.bookstore.ecommerce.app.order.query;

import java.util.List;
import java.util.concurrent.CompletableFuture;
import java.util.stream.Collectors;
import com.bookstore.ecommerce.app.repository.query.BooksQuery;
import com.bookstore.ecommerce.app.repository.query.OrderDetailsQuery;
import org.springframework.stereotype.Component;
import lombok.var;

@Component
public class QueryImpl implements Query {
  private final OrderDetailsQuery orderDetailsQuery;
  private final BooksQuery booksQuery;

  public QueryImpl(OrderDetailsQuery orderDetailsQuery, BooksQuery booksQuery) {
    this.orderDetailsQuery = orderDetailsQuery;
    this.booksQuery = booksQuery;
  }

  @Override
  public CompletableFuture<List<Book>> findBooks(String term) throws Exception {
    var books = this.booksQuery
      .execute(BooksQuery.Params.builder().term(term).build())
      .join();

    return CompletableFuture.completedFuture(
      books
        .stream()
        .map(book -> Book.fromDataObject(book))
        .collect(Collectors.toList()));

  }

  @Override
  public CompletableFuture<Order> getOrderDetails(String orderId) throws Exception {
    final var order = this.orderDetailsQuery
      .execute(OrderDetailsQuery.Params.builder().orderId(orderId).build())
      .join();

    return CompletableFuture.completedFuture(Order.fromDataObject(order));
  }
}
