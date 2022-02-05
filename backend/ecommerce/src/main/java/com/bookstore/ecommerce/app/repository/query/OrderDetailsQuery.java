package com.bookstore.ecommerce.app.repository.query;

import com.bookstore.ecommerce.app.domain.data.Order;
import com.bookstore.ecommerce.app.repository.Query;

import lombok.Builder;
import lombok.Data;

public interface OrderDetailsQuery extends Query<OrderDetailsQuery.Params, Order> {
  @Data
  @Builder
  public class Params {
    String orderId;
  }
}
