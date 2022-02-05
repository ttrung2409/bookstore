package com.bookstore.ecommerce.repository;

import java.util.concurrent.CompletableFuture;
import com.bookstore.ecommerce.app.domain.Order;
import com.bookstore.ecommerce.app.repository.OrderRepository;
import com.bookstore.ecommerce.app.repository.Transaction;
import org.springframework.stereotype.Component;
import lombok.var;

@Component
public class CassandraOrderRepository implements OrderRepository {
  @Override
  public CompletableFuture<Void> create(Order order, Transaction tx) throws Exception {
    var dataOrder = order.getState();

    return CompletableFuture.runAsync(() -> {
      var manager = tx != null ? ((CassandraTransaction) tx).getManager() : new EntityManager();

      try {
        manager.getManager().persist(dataOrder);

        for (var item : dataOrder.getItems()) {
          manager.getManager().persist(item);
        }
      } finally {
        manager.close();
      }
    });
  }
}
