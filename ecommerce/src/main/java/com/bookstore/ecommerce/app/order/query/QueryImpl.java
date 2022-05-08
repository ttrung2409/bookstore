package com.bookstore.ecommerce.app.order.query;

import java.util.concurrent.CompletableFuture;
import com.bookstore.ecommerce.app.repository.queries.OrderDetailsQuery;
import org.springframework.stereotype.Component;
import lombok.var;

@Component
public class QueryImpl implements Query {
  private final OrderDetailsQuery orderDetailsQuery;

  public QueryImpl(OrderDetailsQuery orderDetailsQuery) {
    this.orderDetailsQuery = orderDetailsQuery;
  }

  @Override
  public CompletableFuture<Order> getOrderDetails(String orderId) throws Exception {
    final var order = this.orderDetailsQuery
      .execute(OrderDetailsQuery.Params.builder().orderId(orderId).build())
      .get();

    return CompletableFuture.completedFuture(Order.fromDataObject(order));
  }
}
