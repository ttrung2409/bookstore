package com.bookstore.ecommerce.cassandra;

import com.datastax.oss.driver.api.mapper.annotations.Entity;
import com.datastax.oss.driver.api.mapper.annotations.PartitionKey;
import com.datastax.oss.driver.api.mapper.annotations.PropertyStrategy;

@Entity
@PropertyStrategy(mutable = false)
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
