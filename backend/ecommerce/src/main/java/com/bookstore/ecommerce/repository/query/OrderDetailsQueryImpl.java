package com.bookstore.ecommerce.repository.query;

import java.util.concurrent.CompletableFuture;
import com.bookstore.ecommerce.app.domain.data.Order;
import com.bookstore.ecommerce.app.repository.query.OrderDetailsQuery;
import com.bookstore.ecommerce.repository.EntityManager;
import com.bookstore.ecommerce.utils.NotFoundException;
import lombok.var;

public class OrderDetailsQueryImpl implements OrderDetailsQuery {

  @Override
  public CompletableFuture<Order> execute(Params params) {
    return CompletableFuture.supplyAsync(() -> {
      try (var manager = new EntityManager()) {
        var order = manager.getManager().find(Order.class, params.getOrderId());

        if (order == null) {
          throw new NotFoundException();
        }

        var items = manager.getManager()
          .createQuery(
            String.format("select * from order_item where order_id = \"%s\"", params.getOrderId()))
          .getResultList();

        return order;
      }
    });
  }
}


