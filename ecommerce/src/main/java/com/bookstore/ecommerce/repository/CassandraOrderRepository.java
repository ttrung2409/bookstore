package com.bookstore.ecommerce.repository;

import java.util.concurrent.CompletableFuture;
import com.bookstore.ecommerce.app.domain.Order;
import com.bookstore.ecommerce.app.repository.OrderRepository;
import com.bookstore.ecommerce.app.repository.Transaction;
import com.bookstore.ecommerce.utils.NotFoundException;
import org.springframework.stereotype.Component;
import lombok.var;

@Component
public class CassandraOrderRepository implements OrderRepository {
  @Override
  public CompletableFuture<Void> create(Order order, Transaction tx) throws Exception {
    return CompletableFuture.runAsync(() -> {
      try (var manager = tx != null ? ((CassandraTransaction) tx).getManager()
        : EntityManagerFactory.getInstance().create()) {
        manager.getManager().persist(order.getState());
      }
    });
  }

  @Override
  public CompletableFuture<Order> get(String id, Transaction tx) throws Exception {
    return CompletableFuture.supplyAsync(() -> {
      try (var manager = tx != null ? ((CassandraTransaction) tx).getManager()
        : EntityManagerFactory.getInstance().create()) {
        var order =
          manager.getManager().find(com.bookstore.ecommerce.app.domain.data.Order.class, id);

        if (order == null) {
          throw new NotFoundException();
        }

        return new Order(order);
      }
    });
  }

  @Override
  public CompletableFuture<Void> update(Order order, Transaction tx) throws Exception {
    return CompletableFuture.runAsync(() -> {
      try (var manager = tx != null ? ((CassandraTransaction) tx).getManager()
        : EntityManagerFactory.getInstance().create()) {
        manager.getManager().persist(order.getState());
      }
    });
  }
}
