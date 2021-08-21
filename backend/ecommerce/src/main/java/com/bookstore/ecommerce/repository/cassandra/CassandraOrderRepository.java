package com.bookstore.ecommerce.cassandra;

import java.util.concurrent.CompletableFuture;

import com.bookstore.ecommerce.app.domain.data.Order;
import com.bookstore.ecommerce.app.repository.OrderRepository;
import com.bookstore.ecommerce.app.repository.Transaction;

import org.springframework.stereotype.Component;

@Component
public class CassandraOrderRepository implements OrderRepository {
  @Override
  public CompletableFuture<Order> get(String id) {
    return CompletableFuture.supplyAsync(() -> {
      try (var manager = new EntityManager()) {
        return manager.getManager().find(Order.class, id);
      }
    });
  }

  @Override
  public CompletableFuture<Void> create(Order order, Transaction tx) throws Exception {
    return CompletableFuture.runAsync(() -> {
      var manager = tx != null ? ((CassandraTransaction) tx).getManager() : new EntityManager();

      try {
        manager.getManager().persist(order);

        for (var item : order.getItems()) {
          manager.getManager().persist(item);
        }
      } finally {
        manager.close();
      }
    });
  }
}
