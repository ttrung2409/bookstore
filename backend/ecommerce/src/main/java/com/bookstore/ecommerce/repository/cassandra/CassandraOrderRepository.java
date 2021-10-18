package com.bookstore.ecommerce.repository.cassandra;

import java.util.concurrent.CompletableFuture;
import com.bookstore.ecommerce.app.domain.data.Order;
import com.bookstore.ecommerce.app.repository.OrderRepository;
import com.bookstore.ecommerce.app.repository.Transaction;
import com.bookstore.ecommerce.utils.NotFoundException;
import org.springframework.stereotype.Component;

@Component
public class CassandraOrderRepository implements OrderRepository {
  @Override
  public CompletableFuture<Order> getDetails(String id) throws Exception {
    return CompletableFuture.supplyAsync(() -> {
      try (var manager = new EntityManager()) {
        final var order = manager.getManager().find(Order.class, id);

        if (order == null) {
          throw new NotFoundException();
        }

        final var items = manager.getManager()
          .createQuery(String.format("select * from order_item where order_id = \"%s\"", id))
          .getResultList();

        return order;
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
