package com.bookstore.ecommerce.repository.queries;

import java.util.concurrent.CompletableFuture;
import com.bookstore.ecommerce.app.domain.data.Order;
import com.bookstore.ecommerce.app.repository.queries.OrderDetailsQuery;
import com.bookstore.ecommerce.repository.EntityManagerFactory;
import com.bookstore.ecommerce.utils.NotFoundException;
import lombok.var;

public class OrderDetailsQueryImpl implements OrderDetailsQuery {

  @Override
  public CompletableFuture<Order> execute(Params params) throws Exception {
    return CompletableFuture.supplyAsync(() -> {
      try (var manager = EntityManagerFactory.getInstance().create()) {
        var order = manager.getManager().find(Order.class, params.getOrderId());

        if (order == null) {
          throw new NotFoundException();
        }

        return order;
      }
    });
  }
}
