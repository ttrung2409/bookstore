package com.bookstore.ecommerce.app.order.query;

import java.util.concurrent.CompletableFuture;

public interface Query {
  CompletableFuture<Order> getOrderDetails(String orderId) throws Exception;
}
