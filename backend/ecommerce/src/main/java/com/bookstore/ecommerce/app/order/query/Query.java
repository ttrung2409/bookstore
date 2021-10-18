package com.bookstore.ecommerce.app.order.query;

import java.util.List;
import java.util.concurrent.CompletableFuture;
import com.bookstore.ecommerce.app.order.query.dto.Book;
import com.bookstore.ecommerce.app.order.query.dto.Order;

public interface Query {
  CompletableFuture<List<Book>> findBooks(String term) throws Exception;

  CompletableFuture<Order> getOrderDetails(String orderId) throws Exception;
}
