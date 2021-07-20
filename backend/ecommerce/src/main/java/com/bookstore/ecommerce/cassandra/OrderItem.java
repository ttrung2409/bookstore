package com.bookstore.ecommerce.cassandra;

import com.datastax.oss.driver.api.mapper.annotations.PartitionKey;

public class OrderItem extends com.bookstore.ecommerce.app.data.OrderItem {
  @Override
  @PartitionKey(1)
  public String getOrderId() {
    return super.getOrderId();
  }
  
  @Override
  @PartitionKey(2)
  public String getBookId() {
    return super.getBookId();
  }
}
