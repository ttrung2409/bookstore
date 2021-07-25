package com.bookstore.ecommerce.cassandra;

import java.util.concurrent.CompletableFuture;

import com.bookstore.ecommerce.app.domain.data.Order;
import com.bookstore.ecommerce.app.repository.OrderRepository;

import org.springframework.stereotype.Component;

@Component
public class OrderRepositoryImpl implements OrderRepository {
  @Override
  public CompletableFuture<Order> get(String id) {
    return CompletableFuture.supplyAsync(() -> {
      try (var manager = new EntityManager()) {
        return manager.getManager().find(Order.class, id);
      }
    });
  }

  @Override
  public CompletableFuture<Void> create(Order order) throws Exception {
    return CompletableFuture.runAsync(() -> {
      try (var manager = new EntityManager()) {
        manager.getManager().persist(order);

        for (var item : order.getItems()) {
          manager.getManager().persist(item);
        }
      }
    });
  }
}
