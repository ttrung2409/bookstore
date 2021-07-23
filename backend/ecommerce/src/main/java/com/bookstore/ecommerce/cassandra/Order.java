package com.bookstore.ecommerce.cassandra;

import com.datastax.oss.driver.api.mapper.annotations.Entity;
import com.datastax.oss.driver.api.mapper.annotations.PartitionKey;
import com.datastax.oss.driver.api.mapper.annotations.PropertyStrategy;

@Entity
@PropertyStrategy(mutable = false)
public class Order extends com.bookstore.ecommerce.app.data.Order {
  @Override
  @PartitionKey
  public String getId() {
    return super.getId();
  }
}
