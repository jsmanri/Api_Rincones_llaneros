package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type SitiosTuristicos struct {
	Id                        int         `orm:"column(Id_sitio_turistico);pk;auto"`
	NombreSitioTuristico      string      `orm:"column(Nombre_Sitio_turistico)"`
	DescripcionSitioTuristico string      `orm:"column(Descripcion_Sitio_turistico)"`
	Ubicacion                 string      `orm:"column(Ubicacion)"`
	Horario                   string      `orm:"column(Horario)"`
	IdUsuario                 *Usuarios   `orm:"column(Id_Usuario);rel(fk)"`
	IdCategoria               *Categorias `orm:"column(Id_Categoria);rel(fk)"`
	FotoSitio                 string      `orm:"column(Foto_Sitio);type(text)"`
	FechaCreacion             time.Time   `orm:"column(Fecha_creacion);type(timestamp with time zone);auto_now_add"`
	FechaModifcacion          time.Time   `orm:"column(Fecha_modifcacion);type(timestamp with time zone);auto_now"`
	Activo                    bool        `orm:"column(Activo)"`
}

func (t *SitiosTuristicos) TableName() string {
	return "Sitios_Turisticos"
}

func init() {
	orm.RegisterModel(new(SitiosTuristicos))
}

// AddSitiosTuristicos insert a new SitiosTuristicos into database and returns
// last inserted Id on success.
func AddSitiosTuristicos(m *SitiosTuristicos) (id int64, err error) {
	o := orm.NewOrm()
	if !m.Activo {
		m.Activo = true
	}
	id, err = o.Insert(m)
	return
}

// GetSitiosTuristicosById retrieves SitiosTuristicos by Id. Returns error if
// Id doesn't exist
func GetSitiosTuristicosById(id int) (v *SitiosTuristicos, err error) {
	o := orm.NewOrm()
	v = &SitiosTuristicos{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllSitiosTuristicos retrieves all SitiosTuristicos matches certain condition. Returns empty list if
// no records exist
func GetAllSitiosTuristicos(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SitiosTuristicos)).RelatedSel()
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("error: unused 'order' fields")
		}
	}

	var l []SitiosTuristicos
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateSitiosTuristicos updates SitiosTuristicos by Id and returns error if
// the record to be updated doesn't exist
func UpdateSitiosTuristicosById(m *SitiosTuristicos) (err error) {
	o := orm.NewOrm()
	if !m.Activo {
		m.Activo = true
	}
	v := SitiosTuristicos{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSitiosTuristicos deletes SitiosTuristicos by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSitiosTuristicos(id int) (err error) {
	o := orm.NewOrm()
	v := SitiosTuristicos{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&SitiosTuristicos{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
