package com.bookstore.ecommerce.app.order.command;

import java.util.concurrent.CompletableFuture;

public interface Command {
  CompletableFuture<String> createOrder(CreateOrderRequest request) throws Exception;
}
