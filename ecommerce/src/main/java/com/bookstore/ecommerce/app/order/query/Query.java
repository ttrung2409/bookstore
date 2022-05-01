package com.bookstore.ecommerce.app.order.query;

import java.util.List;
import java.util.concurrent.CompletableFuture;

public interface Query {
  CompletableFuture<List<Book>> findBooks(String term) throws Exception;

  CompletableFuture<Order> getOrderDetails(String orderId) throws Exception;
}
