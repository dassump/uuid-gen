# uuid-gen

uuid-gen is a command line tool to generate UUID Version1 / Version4.

## Getting started

1. Download a pre-compiled binary from the release [page](https://github.com/dassump/uuid-gen/releases).
2. Run `uuid-gen --help`

```shell
$ uuid-gen --help
uuid-gen (dev)

UUID Generator (Version1 / Version4)
https://github.com/dassump/uuid-gen

Usage of uuid-gen:
  -count int
        How many to generate (default 1)
  -quotation string
        Quotation character
  -separator string
        Separator character (default "\n")
  -type string
        UUID Version1 (v1) / Version4 (v4) (default "v4")
```

## Examples

### Version4

```shell
$ uuid-gen
beeb7405-5f81-4de9-a315-9b41578e29c2

$ uuid-gen -type v4
bc7f070d-90ac-4c67-9885-9b867899b418
```

### Version1
```shell
$ uuid-gen -type v1
9cdbde10-2ade-11ed-ab94-fe0d23139169
```

### Count

```shell
$ uuid-gen -count 10
f1073fd6-7c8f-450d-9836-3210381de0b0
4c182baf-3fe3-4b77-b01e-472fcb028f28
9da0f663-4099-4f4f-b385-b74de3815943
1c5d574d-2cd8-4567-9066-ed1e48891bd7
0bc32275-8b3a-4dc2-8b6e-9f39868f0663
2935986a-2191-4f88-867f-4b696651121b
fec7588f-3aa0-4e3d-916f-7df6080ed313
f9a588c1-a58e-4b34-9a3f-90a49b0b0a2d
f40c352e-5bc2-4619-ae58-49b6a777a365
4837bbdc-f9ef-4a26-b912-86c4187d5dd7

$ uuid-gen -type v1 -count 10
b4d05276-2ade-11ed-83a6-fe0d23139169
b4d058a2-2ade-11ed-83a6-fe0d23139169
b4d058ca-2ade-11ed-83a6-fe0d23139169
b4d058e8-2ade-11ed-83a6-fe0d23139169
b4d05906-2ade-11ed-83a6-fe0d23139169
b4d05924-2ade-11ed-83a6-fe0d23139169
b4d05942-2ade-11ed-83a6-fe0d23139169
b4d05960-2ade-11ed-83a6-fe0d23139169
b4d05988-2ade-11ed-83a6-fe0d23139169
b4d059a6-2ade-11ed-83a6-fe0d23139169
```

### Separator

```shell
$ uuid-gen -count 10 -separator ,
19c47af4-7377-4e36-ba5f-c117c6a83227,3746d7c7-13dc-487e-9539-fa193a46c7e7,6cb1f99d-9fda-4108-8811-cda01ea2b202,045500ae-bda4-425a-ad19-5cde083f1d70,c0b0cc4c-f52b-43d4-a8b6-69a8fecdbf8d,45485dc8-7efd-4c43-933f-b7ae1f78c72f,a96d4e01-739e-4b1c-9f2c-7eb1409030e7,bd91b150-5724-4f38-a46c-8648fbaf42ab,8a652d30-dd1d-4fcf-a82b-b2e73ba0e7b6,95b2195b-29b7-43fc-bf01-523a55508b68

$ uuid-gen -count 10 -separator "; "
3a2b712b-10b2-4e95-9684-8f07c82d2648; a935fa37-83d8-4f78-9c21-6890ae873747; 3a49b0fe-c459-4748-9ff1-97e00e3379e0; 2be03423-f4db-4686-8d13-d438cf8268d8; 44507b11-07aa-44e2-bdde-ae9267c30083; 60426747-5a8f-42e9-ae47-7b6652a6c956; bd68985b-4981-4ec4-a770-36ddb239a112; 6254f389-9a62-40e9-b83f-408f4b1a0556; c1133f12-2cfb-46dc-ae99-c94f3db94069; dc0878f2-897d-4cd9-84ad-3424843f9e6d
```

### Quotation

```shell
$ uuid-gen -count 10 -quotation \"
"d3fb7738-5c81-48ac-8a7f-161bd9d7d771"
"6aa9c627-99df-44a4-abd4-39ba3b3858ad"
"930ff5c4-d8f6-4d36-8357-76df624dbc2f"
"3352f2d7-768a-44c5-9279-b623693742dc"
"be44ab5a-a44a-465f-bacd-56d5ce0b468b"
"9756d556-1755-415f-b77a-e52bdae92fb3"
"0dc29886-b96d-451b-84e7-05406f93ca8e"
"1951451e-2041-4b5b-99f2-3acabd079917"
"fd9729f3-a44f-426c-bb07-1def286bfdbd"
"386aee2d-fc69-4b56-acf7-464528036bb4"

$ uuid-gen -count 10 -quotation \'
'6af0a7ef-aac3-483c-8777-2c8e84c4aba3'
'9e67512f-267c-4dd3-b91e-dd1cc55b11c9'
'965c7f2f-3c46-4f5e-944c-2129f73e6250'
'c41bb993-6c58-4138-a0c4-80976a4252b6'
'a21497b5-3933-45f4-bc5d-5708ac4a4a95'
'37cfb4da-09b4-44a2-be9d-f7cf5ba2b06e'
'38f179c3-3662-4ce1-8454-71c1372e91fe'
'292a939e-0576-4117-8d16-24cbea4b69eb'
'05baebfb-0013-40cb-97d0-521d0f6268ff'
'c619c0f1-c4ba-411d-a77f-e269c2c2f42c'
```

### Together

```shell
$ uuid-gen -type v4 -count 10 -quotation \" -separator ,   
"2fe89fc3-0299-419d-a4ec-932d253e55d7","fe0a2126-4c43-489e-a4d1-e1d240ded9c1","c59e0f76-fcc9-4325-a31a-c4882a3fe363","ad8897f0-caf5-47e1-a4c8-cc0afaac21a6","9dea86ff-d0be-4267-ab57-dffe1e6b1c4c","8f8db9af-aef4-443e-a354-a1164af9557d","83124d13-a574-4edb-98dc-5ec40e0015b7","cc4f9970-6a0f-44ce-97f6-2dd42776b77f","40e95a40-79f1-494c-958b-d58a512ed6cf","af3f5de9-711e-4d16-a3f5-46df8dfda7cb"
```

### Shell

```shell
echo "My UUID is `uuid-gen`"
My UUID is becc9377-4262-425b-b1d0-58d9a91f9b6b
```

```shell
export UUIDS=(`uuid-gen -count 10`)

echo $UUIDS[1]
d1dfda37-fa29-4e8d-984a-79b25824a363

echo $UUIDS[2]
44d4b076-bf8b-4aff-859a-ba8783505540

echo $UUIDS[10]
6f266244-12e8-4082-8480-853846e22bbf

for UUID in ${UUIDS[@]}; do echo $UUID; done
d1dfda37-fa29-4e8d-984a-79b25824a363
44d4b076-bf8b-4aff-859a-ba8783505540
a6134978-e00c-466d-a16a-30f87e38b57a
3744f354-28e7-4c6c-a5d1-a3a61e7eb067
2cb14a8e-8d37-4811-97bf-b78669d0b02c
79e78529-46ba-4e1e-af4a-b77234136bd0
f75bb011-9195-432e-87c0-eb9b21924d36
75868796-12f1-47e5-8731-e55664809df3
b3d705c3-ebab-422e-b365-385410e50d24
6f266244-12e8-4082-8480-853846e22bbf
```

```shell
export TMPFILES=()

for x in `uuid-gen -count 10`; do TMPFILES+=`mktemp -q /tmp/tmpfile.$x`; done

for UUID in ${TMPFILES[@]}; do; echo $UUID; done
/tmp/tmpfile.f412389e-c8ff-4113-b14c-d920bd62235b
/tmp/tmpfile.222eb1d1-ba35-4ebc-8aa6-8d29d31fdc09
/tmp/tmpfile.70ca2d1b-2701-4ed0-b483-7b135fe6c957
/tmp/tmpfile.f582ee5c-c2af-480c-a4ba-09c0daa7e143
/tmp/tmpfile.98c9e351-a0d2-443a-b33b-52a569ec57a9
/tmp/tmpfile.9624c23f-fada-4390-9692-af18c92f6446
/tmp/tmpfile.bcf8f64a-f249-4535-aa4c-b14fd0d84497
/tmp/tmpfile.1e9541a4-fb1c-4c4f-bd22-327472b52f9e
/tmp/tmpfile.57e61cdb-1413-4494-8e6e-0315b415e922
/tmp/tmpfile.690c558e-ad05-485f-abaf-eb389460b899
```

## Contributing

Bug reports and pull requests are welcome on GitHub at https://github.com/dassump/uuid-gen.