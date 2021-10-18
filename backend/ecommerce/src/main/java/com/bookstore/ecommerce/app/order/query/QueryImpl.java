package com.bookstore.ecommerce.app.order.query;

import java.util.List;
import java.util.concurrent.CompletableFuture;
import java.util.stream.Collectors;
import com.bookstore.ecommerce.app.order.query.dto.Book;
import com.bookstore.ecommerce.app.order.query.dto.Order;
import com.bookstore.ecommerce.app.repository.BookRepository;
import com.bookstore.ecommerce.app.repository.OrderRepository;

import org.springframework.stereotype.Component;

@Component
public class QueryImpl implements Query {
  private final BookRepository bookRepository;
  private final OrderRepository orderRepository;

  public QueryImpl(BookRepository bookRepository, OrderRepository orderRepository) {
    this.bookRepository = bookRepository;
    this.orderRepository = orderRepository;
  }

  @Override
  public CompletableFuture<List<Book>> findBooks(String term) throws Exception {
    var books = this.bookRepository.find(term).join();

    return CompletableFuture.completedFuture(
      books
        .stream()
        .map(book -> Book.fromDataObject(book))
        .collect(Collectors.toList()));

  }

  @Override
  public CompletableFuture<Order> getOrderDetails(String orderId) throws Exception {
    final var order = this.orderRepository.getDetails(orderId).join();

    return CompletableFuture.completedFuture(Order.fromDataObject(order));
  }
}
